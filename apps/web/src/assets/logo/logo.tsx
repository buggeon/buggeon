import { type SVGProps } from 'react'

export function LogoIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="1em"
      height="1em"
      viewBox="0 0 512 512"
      fill="none"
      {...props}
    >
      <defs>
        <linearGradient id="wingGradient" x1="0%" y1="0%" x2="100%" y2="100%">
          <stop offset="0%" stopColor="#B78CFF" />
          <stop offset="100%" stopColor="#5A32F3" />
        </linearGradient>
        <linearGradient id="headGradient" x1="0%" y1="0%" x2="100%" y2="100%">
          <stop offset="0%" stopColor="#7A58FF" />
          <stop offset="100%" stopColor="#3A1FB5" />
        </linearGradient>
        <filter id="shadow" x="-40%" y="-40%" width="180%" height="180%">
          <feDropShadow dx="0" dy="14" stdDeviation="18" floodColor="#000" floodOpacity=".35" />
        </filter>
        <filter id="glow">
          <feGaussianBlur stdDeviation="35" />
        </filter>
      </defs>

      <circle cx="256" cy="260" r="150" fill="#6E46FF" opacity=".18" filter="url(#glow)" />

      <g filter="url(#shadow)">
        <g stroke="#6338F6" strokeWidth="18" strokeLinecap="round">
          <line x1="140" y1="205" x2="105" y2="180" />
          <line x1="120" y1="275" x2="80" y2="275" />
          <line x1="140" y1="345" x2="105" y2="375" />
          <line x1="372" y1="205" x2="407" y2="180" />
          <line x1="392" y1="275" x2="432" y2="275" />
          <line x1="372" y1="345" x2="407" y2="375" />
        </g>

        <path d="M160 190 A96 96 0 0 1 352 190 L352 240 L160 240 Z" fill="url(#headGradient)" />

        <path d="M210 120 L230 65 L256 105 L282 65 L302 120 L256 135 Z" fill="url(#wingGradient)" />

        <circle cx="230" cy="65" r="8" fill="#B78CFF" />
        <circle cx="256" cy="58" r="8" fill="#B78CFF" />
        <circle cx="282" cy="65" r="8" fill="#B78CFF" />

        <g stroke="#6E46FF" strokeWidth="10" strokeLinecap="round">
          <line x1="185" y1="185" x2="150" y2="150" />
          <line x1="327" y1="185" x2="362" y2="150" />
        </g>

        <circle cx="150" cy="150" r="10" fill="#6E46FF" />
        <circle cx="362" cy="150" r="10" fill="#6E46FF" />

        <path d="M256 210 C180 205 125 250 125 325 C125 395 180 445 256 445 Z" fill="url(#wingGradient)" />
        <path d="M256 210 C332 205 387 250 387 325 C387 395 332 445 256 445 Z" fill="url(#wingGradient)" />

        <rect x="248" y="210" width="16" height="235" rx="8" fill="#2B147D" />

        <circle cx="190" cy="285" r="22" fill="#24125E" />
        <circle cx="190" cy="365" r="22" fill="#24125E" />
        <circle cx="322" cy="285" r="22" fill="#24125E" />
        <circle cx="322" cy="365" r="22" fill="#24125E" />
      </g>
    </svg>
  )
}