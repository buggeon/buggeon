import { type SVGProps } from 'react'

export function RefreshIcon(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            {...props}
        >
            <path d="M4 12A8 8 0 0 1 18.93 8" />
            <path d="M20 12A8 8 0 0 1 5.07 16" />
            <polyline points="14 8 19 8 19 3" />
            <polyline points="10 16 5 16 5 21" />
        </svg>
    )
}