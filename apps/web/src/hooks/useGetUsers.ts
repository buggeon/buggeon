import { useQuery } from '@apollo/client/react'
import { GetProjectsDocument, GetUserDocument, GetUsersDocument } from '../graphql/generated/graphql'

export const useGetUsers = () => {

    const { data, loading, error, refetch } = useQuery(GetUsersDocument)

    return { usersData: data, loading, error, refetch } 
}