import { useQuery } from '@apollo/client/react'
import { GetCardsDocument } from '../graphql/generated/graphql'

export const useGetCards = (boardId: string) => {
  const { data, loading, error } = useQuery(GetCardsDocument, {
    variables: { boardId },
  })
  return { data, loading, error }
}