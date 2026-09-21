import axios, { AxiosError, type InternalAxiosRequestConfig } from "axios";
import { refreshAccessToken } from "../utils/resreshToken";

const api = axios.create({
    baseURL: `http://localhost:9187`,
    withCredentials: true
})

api.interceptors.request.use(   
    (config) => {
        const token = localStorage.getItem("accessToken")

        if(token) {
            config.headers.Authorization = `Bearer ${token}`
        }

        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

api.interceptors.response.use(
    
    (response) => response,
    async (error : AxiosError) => {

        const originalRequest = error.config as InternalAxiosRequestConfig & {
            _retry? : boolean
        }

        if(error.response?.status == 401 && (error.response?.data as any).message == "Invalid or empty accessToken" && !originalRequest._retry) {

            console.log("Обновляем токен...")

            originalRequest._retry = true

            try {
                const newAccessToken = await refreshAccessToken()

                if(newAccessToken) {
                    localStorage.setItem("accessToken", newAccessToken)
                }

                return api(originalRequest)
            }
            catch(e) {
                return Promise.reject(e)
            }

        }

        return Promise.reject(error)
    }
)

export default api