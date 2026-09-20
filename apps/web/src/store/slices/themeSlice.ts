import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { AppearanceSettingsType } from "../../screens/accountSettings/accountSettings";

const initialState : AppearanceSettingsType = {
    accentColor: "blue",
    theme: "dark"
}

const appearanceSlice = createSlice({
    name: 'appearance',
    initialState,
    reducers: {
        setAppearance: (state, action : PayloadAction<AppearanceSettingsType>) => {
            state.accentColor = action.payload.accentColor,
            state.theme = action.payload.theme
        },
        clearAppearance: (state, action : PayloadAction<AppearanceSettingsType>) => {
            state.accentColor = "blue",
            state.theme = "dark"
        },
    }
})

export const { setAppearance, clearAppearance } = appearanceSlice.actions
export default appearanceSlice