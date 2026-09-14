import { decodeJwt } from 'jose'
import type { User } from '../store/types/user.interface'
import api from './api'

class UserApi {

    static async login(login : string, password : string) : Promise<User> {
        try{
            const result = await api.post(`/auth/login`, {
                login: login,
                password: password
            })

            if(result.status == 200) {

                localStorage.setItem("accessToken", result.data.accessToken)

                const userData = decodeJwt(result.data.accessToken)

                return {
                    id: userData.user_id as string,
                    name: userData.user_name as string,
                    login: userData.user_login as string,
                    email: userData.user_email as string,
                    avatarUrl: ""
                }

            }
        }
        catch(e) {
            throw e
        }
    }

    static async regist(login : string, password : string, email : string, name : string) : Promise<User> {

        try{

            const result = await api.post(`/auth/register`, {
                login: login,
                password: password,
                email: email,
                name: name
            })

            if(result.status == 200) {

                localStorage.setItem("accessToken", result.data.accessToken)

                const userData = decodeJwt(result.data.accessToken)

                console.log(userData)

                return {
                    id: userData.user_id as string,
                    name: userData.user_name as string,
                    login: userData.user_login as string,
                    email: userData.user_email as string,
                    avatarUrl: ""
                }

            }

        }
        catch(e) {
            throw new Error(e)
        }

    }

}

export default UserApi