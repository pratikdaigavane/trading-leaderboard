import {Inter as FontSans} from "next/font/google"
import "./globals.css";

import {cn} from "@/lib/utils"
import {ThemeProvider} from "@/components/theme-provider";
import {Toaster} from "@/components/ui/toaster";
import {CSPostHogProvider} from "@/app/providers";
import {Provider} from 'jotai'


const fontSans = FontSans({
    subsets: ["latin"], variable: "--font-sans",
})


export const metadata = {
    title: "Token Tracker", description: "",
};

export default function RootLayout({children}) {
    return (<html lang="en">
    <CSPostHogProvider>
        <body className={cn("min-h-screen bg-background font-sans antialiased", fontSans.variable)}>
        <Provider>
            <ThemeProvider attribute="class" defaultTheme="dark" enableSystem>
                {children}
                <Toaster/>
            </ThemeProvider>
        </Provider>
        </body>
    </CSPostHogProvider>
    </html>);
}
