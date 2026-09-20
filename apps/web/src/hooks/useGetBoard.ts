import { useQuery } from '@apollo/client/react'
import { gql } from '@apollo/client'
import { GetBoardDocument } from '../graphql/generated/graphql'

export const useGetBoard = (id: string, fields : string[]) => {

  const query = gql`
    query GetBoard($id: ID!) {
      board(id: $id) {
        ${fields.join("\n")}
      }
    }
  `

  const { data, loading, error } = useQuery(query as typeof GetBoardDocument, {
    variables: {
        boardId: id
    },
  })
  return { data, loading, error }
}