import { type SVGProps } from "react"

export function UploadIcon(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            viewBox="0 0 32 32"
            fill="currentColor"
            {...props}
        >
            <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M3 22v4c0 .796.316 1.559.879 2.121A2.996 2.996 0 0 0 6 29h20c.796 0 1.559-.316 2.121-.879A2.996 2.996 0 0 0 29 26v-4a1 1 0 0 0-2 0v4a.997.997 0 0 1-1 1H6a.997.997 0 0 1-1-1v-4a1 1 0 0 0-2 0ZM15 6.414V19a1 1 0 0 0 2 0V6.414l2.293 2.293a1 1 0 0 0 1.414-1.414l-4-4a.999.999 0 0 0-1.414 0l-4 4a1 1 0 0 0 1.414 1.414L15 6.414Z"
            />
        </svg>
    )
}