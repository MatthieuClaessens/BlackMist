import { useState } from "react";
import ConnectionNode from "../components/ConnectionNode";
import { StartTor } from "../../wailsjs/go/main/App";

export default function HomePage() {
    const [isActive, setIsActive] = useState<boolean>(false);


    const handleToggle = async () => {
        if (!isActive) {
            try {
                const result = await StartTor();
                console.log(result);
                setIsActive(true);
            } catch (err) {
                console.error("Go error:", err);
            }
        } else {
            setIsActive(false);
        }
    };

    return (
        <main className="h-screen w-screen bg-[#09090f] text-zinc-300 font-sans overflow-hidden flex items-center justify-center p-0 select-none">
            <div className="w-125 h-80 border border-white/5 bg-[#09090f] flex flex-col relative overflow-hidden shadow-[0_0_100px_rgba(0,0,0,0.5)]">
                <ConnectionNode 
                    isActive={isActive} 
                    onToggle={handleToggle} 
                />
            </div>
        </main>
    )
}