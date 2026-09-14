import { type SVGProps } from 'react'

export function BoardOutlineIcon(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            viewBox="0 0 24 24"
            fill="none"
            {...props}
        >
            <g fill="none" stroke="currentColor" strokeWidth="1.5">
                <path d="M2 2h20v20H2z" />
                <path d="M5.331 5.5h5v13h-5zm8.338 0h5V12h-5z" />
            </g>
        </svg>
    )
}