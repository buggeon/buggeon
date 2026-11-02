import { use, useEffect, useRef, useState, type Ref, type RefObject } from "react";

function useVisibleCards(cards, containerRef : RefObject<HTMLDivElement>, cardHeightEstimate = 80, gap = 10) {

    const [visibleCount, setVisibleCount] = useState(cards.length)

    useEffect(() => {
        
        const calculateVisible = () => {

            if(!containerRef.current) return;

            const containerHeight = containerRef.current.clientHeight;

            if(containerHeight == 0) return;

            const moreButtonHeight = 50;

            const avaliableSpace = containerHeight - moreButtonHeight

            let count = Math.floor(avaliableSpace / (cardHeightEstimate + gap))

            count = Math.min(count, cards.length)

            setVisibleCount(count)

        }

    }, [cards, containerRef, cardHeightEstimate, gap])

    return visibleCount;
    
}

export default useVisibleCards