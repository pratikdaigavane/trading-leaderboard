"use client"

import * as React from "react"
import { Moon, Sun } from "lucide-react"
import { useTheme } from "next-themes"

import { Button } from "@/components/ui/button"
export function ThemeToggle() {
    const { theme, setTheme } = useTheme()

    return (
        <Button
            variant="ghost"
            size="icon"
            onClick={() => {
                const newTheme = theme === "light" ? "dark" : "light"
                setTheme(newTheme)
            }
        }
        >
            <Sun className="h-8 w-8 dark:hidden" />
            <Moon className="hidden h-8 w-8 dark:block" />
            <span className="sr-only">Toggle theme</span>
        </Button>
    )
}