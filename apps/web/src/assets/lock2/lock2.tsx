import { type SVGProps } from "react";

export function Lock2Icon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      viewBox="0 0 24 24"
      width="1em"
      height="1em"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <defs>
        <mask id="lock-hole">
          <rect width="24" height="24" fill="white" />
          <circle cx="12" cy="14.9" r="1.4" fill="black" />
        </mask>
      </defs>

      <g fill="currentColor" mask="url(#lock-hole)">
        <path d="M12 2.75a3.75 3.75 0 0 0-3.75 3.75v2h1.5v-2a2.25 2.25 0 1 1 4.5 0v2h1.5v-2A3.75 3.75 0 0 0 12 2.75Z" />

        <rect
          x="4.75"
          y="8.5"
          width="14.5"
          height="12.75"
          rx="3"
        />
      </g>
    </svg>
  );
}