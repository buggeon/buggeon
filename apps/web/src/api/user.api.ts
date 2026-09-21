import { decodeJwt } from 'jose'
import type { User } from '../store/types/user.interface'
import api from './api'

class UserApi {

    async login(login : string, password : string) : Promise<User> {
        try{
            const result = await api.post(`/auth/login`, {
                login: login,
                password: password
            })

            if(result.status == 200) {

                console.log(result)

                localStorage.setItem("accessToken", result.data.accessToken)

                const userId = decodeJwt(result.data.accessToken).userId

                return {
                    id: userId as string,
                    name: result.data.userData.name as string,
                    login: result.data.userData.login as string,
                    email: result.data.userData.email as string,
                    avatarUrl: result.data.userData.avatarUrl as string
                }

            }
        }
        catch(e) {
            throw e
        }
    }

    async regist(login : string, password : string, email : string, name : string) : Promise<User> {

        try{

            const result = await api.post(`/auth/register`, {
                login: login,
                password: password,
                email: email,
                name: name
            })

            if(result.status == 200) {

                console.log(result)

                localStorage.setItem("accessToken", result.data.accessToken)

                const userId = decodeJwt(result.data.accessToken).userId

                return {
                    id: userId as string,
                    name: result.data.userData.name,
                    login: result.data.userData.login,
                    email: result.data.userData.email,
                    avatarUrl: result.data.userData.avatarUrl
                }

            }

        }
        catch(e) {
            throw new Error(e)
        }

    }

    async setAvatar(userId : string, avatar : File) : Promise<string> {

        const formData = new FormData()
        formData.append("avatar", avatar)

        try{

            const response = await api.patch(`/users/${userId}/avatar`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })

            return response.data.avatarUrl

        }
        catch(e) {
            throw new Error(e)
        }

    }

}

const userApi = new UserApi()

export default userApi