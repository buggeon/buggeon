import { useQuery } from '@apollo/client/react'
import { GetMembersDocument } from '../graphql/generated/graphql'

export const useGetMembers = (projectId: string) => {
  const { data, loading, error } = useQuery(GetMembersDocument, {
    variables: { projectId },
  })
  return { data, loading, error }
}