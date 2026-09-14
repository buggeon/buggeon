import { useQuery } from '@apollo/client/react'
import { GetProjectsDocument } from '../graphql/generated/graphql'

export const useGetProjects = (userId: string) => {

    const { data, loading, error, refetch } = useQuery(GetProjectsDocument, {
        variables: { userId },
        skip: false,
        fetchPolicy: 'network-only',
    })

    return { data, loading, error, refetch } 
}