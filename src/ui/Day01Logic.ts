import { RefObject, useCallback, useEffect, useRef, useState } from 'react';

type Day01Logic = {
    numbers: number[];
    refTrack: RefObject<any>;
    play: () => void;
    next: () => void;
    reset: () => void;
    state: number;
    sum: number;
    run: boolean;
};

const base = 175;
const numberWidth = 100;

export const useDay01 = (numbers: number[]): Day01Logic => {
    const [index, setIndex] = useState(0);
    const [state, setState] = useState(0);
    const [sum, setSum] = useState(0);
    const [run, setRun] = useState(false);
    const refTrack = useRef<any>();

    // configure children on load
    useEffect(() => {
        const numbers = refTrack.current.children;
        Array.from(numbers).forEach((number: any, i: number) => {
            number.style.left = base + i * numberWidth + 'px';
        });
        refTrack.current.style.width = numberWidth * (numbers.length - 1);
    }, []);

    const move = useCallback(() => {
        if (refTrack && refTrack.current)
            refTrack.current.style.transform = 'translateX(-' + (50 + (index + 1) * numberWidth) + 'px)';
    }, [index]);

    useEffect(() => {
        setState(numbers[index] - numbers[index - 1]);
        setSum((s) => (numbers[index] > numbers[index - 1] ? s + 1 : s));
        move();
    }, [index, move, numbers]);

    const next = useCallback(() => {
        setIndex((n) => (index < numbers.length - 1 ? n + 1 : n));
    }, [index, numbers]);

    const reset = useCallback(() => {
        setIndex(0);
        setState(0);
        setSum(0);
        setRun(false);
        move();
    }, [move]);

    const play = useCallback(() => {
        setRun(!run);
    }, [run]);

    useEffect(() => {
        let timeout: ReturnType<typeof setTimeout>;
        const startTimer = () => {
            timeout = setTimeout(() => next(), 0.0001);
        };
        run && startTimer();
        return () => {
            clearTimeout(timeout);
        };
    }, [next, run]);

    return { numbers, refTrack, play, next, reset, state, sum, run };
};
