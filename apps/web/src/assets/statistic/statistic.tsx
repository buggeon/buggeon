import { type SVGProps } from 'react'

export function StatisticUpIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="1em"
      height="1em"
      fill="none"
      {...props}
    >
      <g
        stroke="currentColor"
        strokeWidth={2}
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <path
          strokeMiterlimit={5.759}
          d="M3 3v16a2 2 0 0 0 2 2h16"
        />
        <path
          strokeMiterlimit={5.759}
          d="m7 14 4-4 4 4 6-6"
        />
        <path d="M18 8h3v3" />
      </g>
    </svg>
  )
}