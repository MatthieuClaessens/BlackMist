import { useState, useEffect } from "react";
import ConnectionNode from "../components/ConnectionNode";
import { TorService } from "../services/TorService";

export default function HomePage() {
    const [isActive, setIsActive] = useState<boolean>(false);
    const [currentIp, setCurrentIp] = useState<string>("---.---.---.---");
    const [isConnecting, setIsConnecting] = useState<boolean>(false);
    const [ping, setPing] = useState<number>(0);


useEffect(() => {
        let interval: ReturnType<typeof setInterval> | undefined;

        const fetchStats = async (): Promise<void> => {
            const [newIp, newPing]: [string, number] = await Promise.all([
                TorService.getIP(),
                TorService.getPing()
            ]);
            setCurrentIp(newIp);
            setPing(newPing);
        };

        if (isActive) {
            fetchStats();
            interval = setInterval(fetchStats, 5000);
        } else {
            setCurrentIp("---.---.---.---");
            setPing(0);
        }

        return () => {
            if (interval) clearInterval(interval);
        };
    }, [isActive]);

const handleToggle = async () => {
        setIsConnecting(true);
        try {
            if (!isActive) {
                await TorService.connect();
                setIsActive(true);
            } else {
                await TorService.disconnect();
                setIsActive(false);
            }
        } catch (err) {
            console.error("Action failed:", err);
        } finally {
            setIsConnecting(false);
        }
    };
        

        return (
            <main className="h-screen w-screen bg-[#09090f] text-zinc-300 font-sans overflow-hidden flex items-center justify-center p-0 select-none">
                <div className="w-125 h-80 border border-white/5 bg-[#09090f] flex flex-col relative overflow-hidden shadow-[0_0_100px_rgba(0,0,0,0.5)]">
                    <ConnectionNode
                        isActive={isActive}
                        onToggle={handleToggle}
                        currentIp={currentIp}
                        isConnecting={isConnecting}
                        ping={ping}
                    />
                </div>
            </main>
        )
    }
