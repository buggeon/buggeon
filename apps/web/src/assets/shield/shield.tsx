import { type SVGProps } from 'react'

export function ShieldIcon(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            viewBox="0 0 24 24"
            fill="currentColor"
            {...props}
        >
            <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M12 2L4 5v6c0 5.55 3.42 10.74 8 12 4.58-1.26 8-6.45 8-12V5l-8-3zm0 2.18l6 2.25V11c0 4.52-2.55 8.85-6 10-3.45-1.15-6-5.48-6-10V6.43l6-2.25z"
            />
        </svg>
    )
}