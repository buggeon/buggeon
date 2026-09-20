import { useContext, useEffect } from "react";
import { useSelector } from "react-redux";
import type { RootState } from "../store/slices";

function ThemeProvider({children}) {

    const appearance  = useSelector((state : RootState) => state.appearance)

    useEffect(() => {

        const root = document.documentElement

        const effectiveTheme =
            appearance.theme === "system"
                ? window.matchMedia("(prefers-color-scheme: dark").matches
                    ? "dark"
                    : "light"
                : appearance.theme

        root.setAttribute('data-theme', effectiveTheme)
        root.setAttribute('data-accent', appearance.accentColor)

    }, [appearance])

    useEffect(() => {

        if(appearance.theme != "system") return

        const mq = window.matchMedia('(prefers-color-scheme: dark')

        const handler = (e) => {
            document.documentElement.setAttribute('data-theme', e.matches ? "dark" : "light")
        }

        mq.addEventListener('change', handler)
        return () => mq.removeEventListener('change', handler)

    }, [appearance.theme])

    return children

}

export default ThemeProvider