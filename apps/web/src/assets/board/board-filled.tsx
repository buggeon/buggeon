import { type SVGProps } from 'react'

export function BoardFilledIcon(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            viewBox="0 0 24 24"
            fill="currentColor"
            {...props}
        >
            <path fillRule="evenodd" d="M23 1H1v22h22zM5 5h5.6v14H5zm14 0h-5.6v7H19z" clipRule="evenodd" />
        </svg>
    )
}