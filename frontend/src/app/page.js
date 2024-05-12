import {ThemeToggle} from "@/components/theme-toggle";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import Image from "next/image";
import CommandMenu from "@/components/command-bar";
import Leaderboard from "@/components/leaderboard";


async function fetchSymbolsList() {
    const res = await fetch('http://localhost:8080/symbols')
    if (res.status === 200) {
        const json = await res.json()
        console.log(json)
        return json
    }
}

export default async function Home() {
    const symbolsList = await fetchSymbolsList()
    return (
        <main className="flex-col h-screen md:flex p-12">
            <header className="flex items-center gap-4 place-content-between">
                <div>
                    <h1 className="text-4xl font-bold">Welcome Back!</h1>
                    <h3 className="text-l text-muted-foreground pt-2">Here is the leaderboard of your favourite
                        token </h3>
                </div>
                <ThemeToggle/>
            </header>
            <div className="mt-8 md:mx-24" >
                <Tabs defaultValue="Bitcoin">
                    <div className="flex items-center gap-4">
                        <CommandMenu dbList={symbolsList}></CommandMenu>
                    </div>
                    <div className="mt-4">
                        <Leaderboard></Leaderboard>
                    </div>
                </Tabs>

            </div>
            <div className="mt-4 pb-3 flex items-center justify-center text-muted-foreground text-sm">
                Made with ❤️ by <div>&nbsp;</div>  <a href="https://www.linkedin.com/in/pratikdaigavane/" target="_blank"> Pratik Daigavane</a>
            </div>
        </main>
    );
}
