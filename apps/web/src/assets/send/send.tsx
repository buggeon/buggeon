import { type SVGProps } from 'react'

export function SendIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="1em"
      height="1em"
      viewBox="0 0 32 32"
      fill="none"
      {...props}
    >
      <path
        fill="currentColor"
        fillRule="evenodd"
        clipRule="evenodd"
        d="M11.499,19.173 L17.3,13.324 C17.689,12.932 18.322,12.93 18.714,13.318 C19.106,13.707 19.109,14.34 18.72,14.732 L12.922,20.579 L18.228,28.581 C18.435,28.894 18.8,29.064 19.173,29.022 C19.546,28.98 19.864,28.733 19.997,28.382 L29.021,4.478 C29.159,4.112 29.071,3.698 28.795,3.42 C28.519,3.142 28.106,3.051 27.738,3.187 L3.734,12.079 C3.381,12.209 3.132,12.527 3.088,12.9 C3.044,13.273 3.213,13.64 3.526,13.848 L11.499,19.173 Z"
      />
    </svg>
  )
}