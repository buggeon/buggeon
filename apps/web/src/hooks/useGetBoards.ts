import { useQuery } from '@apollo/client/react'
import { GetBoardsDocument } from '../graphql/generated/graphql'

export const useGetBoards = (projectId: string) => {

    const { data, loading, error, refetch } = useQuery(GetBoardsDocument, {
        variables: { projectId },
        skip: false
    })

    return { data, loading, error, refetch } 
}