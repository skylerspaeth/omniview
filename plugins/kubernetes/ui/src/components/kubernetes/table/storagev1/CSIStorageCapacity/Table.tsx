import React from 'react'
import { ColumnDef } from '@tanstack/react-table'
import { CSIStorageCapacity } from 'kubernetes-types/storage/v1'
import { useParams } from 'react-router-dom'
import ResourceTable from '../../../../shared/table/ResourceTable'
import { withClusterResourceColumns } from '../../shared/columns'

const resourceKey = 'storage::v1::CSIStorageCapacity'

const CSIStorageCapacityTable: React.FC = () => {
  const { id = '' } = useParams<{ id: string }>()

  const columns = React.useMemo<Array<ColumnDef<CSIStorageCapacity>>>(
    () => withClusterResourceColumns([], { connectionID: id, resourceKey }),
    [],
  )

  return (
    <ResourceTable
      columns={columns}
      connectionID={id}
      resourceKey={resourceKey}
      idAccessor='metadata.uid'
      memoizer='metadata.uid,metadata.resourceVersion'
    />
  )
}

export default CSIStorageCapacityTable
