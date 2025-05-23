import React from 'react';

// ensure we preload deps first before anything
import { preloadSharedDeps } from './features/plugins/api/preloader';
preloadSharedDeps()

import { Provider } from 'react-redux';
import { store } from './store/store';

import './utils/globalutils';
import './providers/monaco/bootstrap';

import {
  experimental_extendTheme as materialExtendTheme,
  Experimental_CssVarsProvider as MaterialCssVarsProvider,
  THEME_ID as MATERIAL_THEME_ID,
} from '@mui/material/styles';

// Providers
import { CssVarsProvider, StyledEngineProvider } from '@mui/joy/styles';
import theme from './theme';
import { AppSnackbarProvider } from '@/contexts/AppSnackbarProvider';
import RightDrawerProvider from '@/providers/RightDrawerProvider';
import { ErrorBoundary } from 'react-error-boundary';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';
import { ExtensionProvider } from '@omniviewdev/runtime';
import BottomDrawerProvider from './providers/BottomDrawer/provider';
import { SettingsProvider } from '@omniviewdev/runtime';
import { EXTENSION_REGISTRY } from './features/extensions/store'
import { ConfirmationModalProvider } from './contexts/ConfirmationModalContext';
import { PluginRegistryProvider } from './features/plugins/PluginRegistryProvider';
import { RouteProvider } from './features/router/RouteProvider';
import log from '@/features/logger'

// 2. Global handlers for uncaught errors + promise rejections
window.addEventListener("error", event => {
  // event.error may be undefined (e.g. script load failures)
  const err = event.error || new Error(String(event.message));
  log.error(err, {
    filename: event.filename,
    lineno: event.lineno,
    colno: event.colno,
  });
});

window.addEventListener("unhandledrejection", event => {
  const reason = event.reason instanceof Error
    ? event.reason
    : new Error(JSON.stringify(event.reason));
  log.error(reason, { unhandledRejection: true });
});

// Create a client
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

function fallbackRender({ error }: { error: Error }) {
  return (
    <div role='alert'>
      <p>Something went wrong:</p>
      <pre style={{ color: 'red' }}>{error.message}</pre>
    </div>
  );
}

const materialTheme = materialExtendTheme();

log.debug("starting up the application")

/**
 * Render out the core layout for the application
 */
const App: React.FC = () => (
  <ErrorBoundary FallbackComponent={fallbackRender}>
    <SettingsProvider>
      <ExtensionProvider registry={EXTENSION_REGISTRY}>
        <AppSnackbarProvider>
          <QueryClientProvider client={queryClient}>
            <MaterialCssVarsProvider
              defaultMode="dark"
              theme={{ [MATERIAL_THEME_ID]: materialTheme }}
            >
              <StyledEngineProvider injectFirst>
                <Provider store={store}>
                  <CssVarsProvider
                    defaultMode='dark'
                    modeStorageKey='omniview_identify-system-mode'
                    // Set as root provider
                    disableNestedContext
                    theme={theme}
                  >
                    <ConfirmationModalProvider>
                      <RightDrawerProvider>
                        <BottomDrawerProvider>
                          <PluginRegistryProvider>
                            <RouteProvider />
                          </PluginRegistryProvider>
                        </BottomDrawerProvider>
                      </RightDrawerProvider>
                    </ConfirmationModalProvider>
                  </CssVarsProvider>
                </Provider>
              </StyledEngineProvider>
            </MaterialCssVarsProvider>
          </QueryClientProvider>
        </AppSnackbarProvider>
      </ExtensionProvider>
    </SettingsProvider>
  </ErrorBoundary>
);

export default App;
