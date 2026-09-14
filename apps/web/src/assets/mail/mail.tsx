import { type SVGProps } from 'react'

export function MailIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="1em"
      height="1em"
      viewBox="0 0 24 24"
      fill="none"
      {...props}
    >
      <path
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
        d="M2.357 7.714L9.337 12.368C10.3 13.009 10.781 13.33 11.301 13.455C11.761 13.565 12.24 13.565 12.7 13.455C13.22 13.33 13.701 13.009 14.664 12.368L21.644 7.714"
      />
      <path
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
        d="M7.157 19.5H16.843C18.523 19.5 19.363 19.5 20.005 19.173C20.586 18.875 21.049 18.412 21.347 17.831C21.674 17.189 21.674 16.349 21.674 14.669V9.3C21.674 7.62 21.674 6.78 21.347 6.138C21.049 5.557 20.586 5.094 20.005 4.796C19.363 4.469 18.523 4.469 16.843 4.469H7.157C5.477 4.469 4.637 4.469 3.995 4.796C3.414 5.094 2.951 5.557 2.653 6.138C2.326 6.78 2.326 7.62 2.326 9.3V14.669C2.326 16.349 2.326 17.189 2.653 17.831C2.951 18.412 3.414 18.875 3.995 19.173C4.637 19.5 5.477 19.5 7.157 19.5Z"
      />
    </svg>
  )
}