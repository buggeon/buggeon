import { useQuery } from '@apollo/client/react'
import { gql } from '@apollo/client'
import { GetCardDocument } from '../graphql/generated/graphql'

export const useGetCard = (id: string, fields : string[]) => {

  const query = gql`
    query GetCard($id: ID!) {
      card(id: $id) {
        ${fields.join("\n")}
      }
    }
  `

  const { data, loading, error } = useQuery(query as typeof GetCardDocument, {
    variables: { id },
  })
  return { data, loading, error }
}