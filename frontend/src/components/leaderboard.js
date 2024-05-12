'use client'

import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import Image from "next/image";
import {useEffect, useRef, useState} from "react";
import {useToast} from "@/components/ui/use-toast"
import {Skeleton} from "@/components/ui/skeleton";
import {symbolImageInstance, symbolInstance, symbolNameInstance} from "@/app/store";
import {useAtom} from "jotai";

function useInterval(callback, delay) {
    const savedCallback = useRef();

    // Remember the latest callback.
    useEffect(() => {
        savedCallback.current = callback;
    }, [callback]);

    // Set up the interval.
    useEffect(() => {
        let id = setInterval(() => {
            savedCallback.current();
        }, delay);
        return () => clearInterval(id);
    }, [delay]);
}


export default function Leaderboard() {
    const [symbol, setSymbol] = useAtom(symbolInstance);
    const [symbolName, setSymbolName] = useAtom(symbolNameInstance);
    const [symbolImage, setSymbolImage] = useAtom(symbolImageInstance);
    const [leaderboard, setLeaderboard] = useState([])
    const [counter, setCounter] = useState(60)
    const {toast} = useToast()
    const [loading, setLoading] = useState(true)

    async function fetchLeaderboard() {
        if (counter === 60) {
            setLoading(true)
        }
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_API_URL}/leaderboard/${symbol}`)
        if (res.status === 200) {
            const json = await res.json()
            console.log(json)
            setLeaderboard(json)
            setLoading(false)
            return json
        }
    }

    useEffect(() => {
            console.log("fetching leaderboard", symbol)
            fetchLeaderboard()
        },
        [symbol])
    useInterval(() => {
            setCounter(counter - 1)
            console.log("sdfsdf", counter)
            if (counter === 0) {
                fetchLeaderboard()
                setCounter(60)
                toast({
                    title: "Leaderboard refreshed",
                })
            } else if (counter === 60) {
                fetchLeaderboard()
            } else {
                setCounter(counter - 1)
                console.log("counter is", counter - 1, " symbol is ", symbol )
            }
        },
        1000);


    // useEffect(() => {
    //     setInterval(
    //     }, 1000)
    // }, []);

    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    <div className="flex items-center gap-4">
                        <Image
                            alt="Product image"
                            className="aspect-square rounded-md object-cover"
                            height="64"
                            src={symbolImage}
                            width="64"
                        />
                        <div className="flex-col space-y-2">
                            <div> {symbolName}</div>
                            <CardDescription>
                                A list of top 10 traders in the last 24 hours
                            </CardDescription>
                        </div>
                    </div>

                </CardTitle>

            </CardHeader>
            <CardContent>
                {loading === true ?
                    <>
                        <div className="flex-col gap-8">
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                            <div className="flex items-center space-x-4 mt-8">
                                <Skeleton className="h-12 w-12 rounded-full"/>
                                <div className="space-y-2">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden lg:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                <div className="space-y-2 hidden 2xl:block">
                                    <Skeleton className="h-4 w-[250px]"/>
                                    <Skeleton className="h-4 w-[200px]"/>
                                </div>
                                {/*<div className="space-y-2 hidden 2xl:block">*/}
                                {/*    <Skeleton className="h-4 w-[250px]"/>*/}
                                {/*    <Skeleton className="h-4 w-[200px]"/>*/}
                                {/*</div>*/}
                            </div>
                        </div>
                    </> :
                    <>
                        <Table>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>#</TableHead>
                                    <TableHead>Avatar</TableHead>
                                    <TableHead>User</TableHead>
                                    <TableHead>Trading volume(24h)</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {
                                    leaderboard.map((user, index) => {
                                        return (
                                            <TableRow key={user.traderId}>
                                                <TableCell>
                                                    {user.rank}
                                                </TableCell>
                                                <TableCell>
                                                    <Image
                                                        alt="Product image"
                                                        className="aspect-square rounded-md object-cover"
                                                        height="64"
                                                        src={`https://robohash.org/${user.traderId}.png`}
                                                        width="64"
                                                    />
                                                </TableCell>
                                                <TableCell className="font-medium">
                                                    {user.traderId}
                                                </TableCell>
                                                <TableCell>
                                                    {user.totalVolume.toFixed(2)}
                                                </TableCell>
                                            </TableRow>
                                        )
                                    })
                                }
                            </TableBody>
                        </Table>
                    </>}

            </CardContent>
            <CardFooter>
                <div className="text-xs text-muted-foreground font-mono">
                    {
                        counter === 0 ? ("Refreshing now") : (`Refreshing leaderboard in ${counter} seconds`)
                    }

                </div>
            </CardFooter>
        </Card>
    )
}