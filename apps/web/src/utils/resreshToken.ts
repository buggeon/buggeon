import axios from 'axios'
import { API_URL } from '../../config'

let isRefreshing = false
let failedQueue: Array<{
    resolve: (value: any) => void
    reject: (reason?: any) => void
}> = []

const processQueue = (error: Error | null, token: string | null = null) => {
    failedQueue.forEach((promise) => {
        if (error) {
            promise.reject(error)
        } else {
            promise.resolve(token)
        }
    })
    failedQueue = []
}

export const refreshAccessToken = (): Promise<string> => {
    return new Promise((resolve, reject) => {
        if (isRefreshing) {
            failedQueue.push({ resolve, reject })
            return
        }

        isRefreshing = true

        axios.post(
            `${API_URL}/auth/refreshtoken`,
            {},
            { withCredentials: true }
        )
            .then((response) => {
                const newAccessToken = response.data.accessToken
                if (newAccessToken) {
                    localStorage.setItem('accessToken', newAccessToken)
                }
                processQueue(null, newAccessToken)
                isRefreshing = false
                resolve(newAccessToken)
            })
            .catch((error) => {
                console.log("Ошибка при обновлении токена: " + error)
                processQueue(error, null)
                isRefreshing = false
                localStorage.removeItem('accessToken')
                reject(error)
            })
    })
}