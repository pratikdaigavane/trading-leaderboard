'use client'

import React, {useState, useEffect} from "react";
import {
    Calculator,
    Calendar,
    CreditCard, DatabaseZap,
    Settings,
    Smile,
    User,
} from "lucide-react"

import {
    CommandDialog,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    CommandSeparator,
    CommandShortcut,
} from "@/components/ui/command"
import {Button} from "@/components/ui/button";
import { ChevronDown } from "lucide-react"
import Image from "next/image";
import {useAtom} from "jotai";
import {symbolCodeInstance, symbolImageInstance, symbolInstance, symbolNameInstance} from "@/app/store";

export default function CommandMenu({dbList}) {
    const [open, setOpen] = useState(false)
    const [symbol, setSymbol] = useAtom(symbolInstance);
    const [symbolName, setSymbolName] = useAtom(symbolNameInstance);
    const [symbolImage, setSymbolImage] = useAtom(symbolImageInstance);
    const [symbolCode, setSymbolCode] = useAtom(symbolCodeInstance);

    useEffect(() => {
        const down = (e) => {
            if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
                e.preventDefault()
                setOpen((open) => !open)
            }
        }

        document.addEventListener("keydown", down)
        return () => document.removeEventListener("keydown", down)
    }, [])


    function handleSelect(value){
        setSymbol(value.id)
        setSymbolName(value.name)
        setSymbolImage(value.imagePath)
        setSymbolCode(value.symbol)
        setOpen(false)
    }

    return <>
        <div>
            <Button onClick={() => setOpen(true)}
                    className="inline-flex items-center rounded-md font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 border border-input bg-transparent shadow-sm hover:bg-accent hover:text-accent-foreground h-10 px-2 py-2 relative w-full justify-start text-sm text-muted-foreground sm:pr-12 md:w-[20rem] lg:w-[30rem]">
                <Image
                    alt="Product image"
                    className="aspect-square rounded-md object-cover mr-2"
                    height="16"
                    src={symbolImage}
                    width="16"
                />
                <span className="hidden lg:inline-flex">{symbolName}</span>
                <span className="inline-flex lg:hidden">Search...</span>
                <span className="pointer-events-none absolute right-2 top-2 hidden h-5 select-none items-center gap-1 px-1.5 opacity-100 sm:flex">
                   <ChevronDown />
                </span>
            </Button>

        </div>
        <CommandDialog open={open} onOpenChange={setOpen}>
            <CommandInput placeholder="Type a token or search..."/>
            <CommandList>
                <CommandEmpty>No results found.</CommandEmpty>
                <CommandGroup heading="Tokens">
                    {
                        Object.entries(dbList).map(([key, value]) => {
                            return (
                                <CommandItem key={key} onSelect={() => handleSelect(value)} >
                                        <Image
                                            alt="Product image"
                                            className="aspect-square rounded-md object-cover mr-2"
                                            height="32"
                                            src={value.imagePath}
                                            width="32"
                                        />
                                        <span>{value.name}</span>
                                </CommandItem>
                            )
                        })
                    }
                </CommandGroup>
            </CommandList>
        </CommandDialog>
    </>
}