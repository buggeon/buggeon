import type { User } from '../store/types/user.interface'
import api from './api'

class SystemApi {

    static async getAllUsers() : Promise<User[]> {
        try{
            const result = (await api.get("/users"))

            if(result.status == 200) {

                let users = result.data.map(user => ({
                    id: user.id as string,
                    name: user.name as string,
                    login: user.login as string,
                    email: user.email as string,
                    avatarUrl: ""
                }))

                return users

            }
        }
        catch(e) {
            throw new Error(e)
        }
    }

}

export default SystemApi