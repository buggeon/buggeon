import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { AppearanceSettingsType } from "../../screens/accountSettings/accountSettings";

const initialState : AppearanceSettingsType = {
    accentColor: "#4da592",
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
            state.accentColor = "#4da592",
            state.theme = "dark"
        },
    }
})

export const { setAppearance, clearAppearance } = appearanceSlice.actions
export default appearanceSlice