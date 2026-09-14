import { useQuery } from '@apollo/client/react'
import { gql } from '@apollo/client'
import { GetProjectDocument } from '../graphql/generated/graphql'

export const useGetProject = (id: string, fields : string[]) => {

  const query = gql`
    query GetProject($id: ID!) {
      project(id: $id) {
        ${fields.join("\n")}
      }
    }
  `

  const { data, loading, error } = useQuery(query as typeof GetProjectDocument, {
    variables: { id },
  })
  return { data, loading, error }
}