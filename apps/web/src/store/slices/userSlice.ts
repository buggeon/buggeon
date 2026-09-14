import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { User } from "../types/user.interface";

const initialState : User = {
    name: "",
    email: "",
    login: "",
    id: "",
    avatarUrl: ""
}

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action : PayloadAction<User>) => {
            state.name = action.payload.name,
            state.id = action.payload.id,
            state.login = action.payload.login,
            state.email = action.payload.email,
            state.avatarUrl = action.payload.avatarUrl
        },
        clearUser: (state, action : PayloadAction<User>) => {
            state.name = "",
            state.id = "",
            state.login = "",
            state.avatarUrl = "",
            state.email = ""
        },
    }
})

export const { setUser, clearUser } = userSlice.actions
export default userSlice