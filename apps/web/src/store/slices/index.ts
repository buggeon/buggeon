import { combineReducers, configureStore } from "@reduxjs/toolkit";
import userSlice from "./userSlice";
import { FLUSH, PAUSE, PERSIST, persistReducer, persistStore, PURGE, REGISTER, REHYDRATE } from "redux-persist";
import appearanceSlice from "./themeSlice";

const storage = {
    getItem: (key: string) => {
        try {
            const value = localStorage.getItem(key);
            return Promise.resolve(value ? JSON.parse(value) : null);
        } catch (error) {
            return Promise.resolve(null);
        }
    },
    setItem: (key: string, value: any) => {
        try {
            localStorage.setItem(key, JSON.stringify(value));
            return Promise.resolve(value);
        } catch (error) {
            return Promise.resolve(null);
        }
    },
    removeItem: (key: string) => {
        try {
            localStorage.removeItem(key);
            return Promise.resolve();
        } catch (error) {
            return Promise.resolve();
        }
    },
};

const persistConfig = {
    key: "root",
    storage,
    whitelist: ['user', "appearance"],
    blacklist: []
}

const rootReducer = combineReducers({
    user: userSlice.reducer,
    appearance: appearanceSlice.reducer
})

const persistedReducer = persistReducer(persistConfig, rootReducer)

export const store = configureStore({
    reducer: persistedReducer,
    middleware: (getDefaultMiddleware) => 
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
            }
        })
})

export const persistor = persistStore(store)

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch