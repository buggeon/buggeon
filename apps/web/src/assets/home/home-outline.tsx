import { type SVGProps } from 'react'

export function HomeOutlineIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="1em"
      height="1em"
      viewBox="0 0 16 16"
      fill="none"
      {...props}
    >
      <path
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinejoin="round"
        d="M1.5 6L8 1L14.5 6V14.5H11V11C11 9.34315 9.65685 8 8 8C6.34315 8 5 9.34315 5 11V14.5H1.5V6Z"
      />
    </svg>
  )
}