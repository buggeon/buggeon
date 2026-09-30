import { useQuery } from '@apollo/client/react'
import { gql } from '@apollo/client'
import { GetCardDocument, GetMemberDocument } from '../graphql/generated/graphql'

export const useGetMember = (id: string, fields : string[]) => {

  const query = gql`
    query GetMember($memberId: ID!) {
      member(memberId: $memberId) {
        ${fields.join("\n")}
      }
    }
  `

  const { data, loading, error } = useQuery(query as typeof GetMemberDocument, {
    variables: { memberId: id },
  })
  return { memberData: data, loading, error }
}