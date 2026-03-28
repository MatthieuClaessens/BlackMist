import { ShieldCheck, Power, Wifi } from "lucide-react";


interface ConnectionNodeProps {
    isActive: boolean;
    onToggle: () => void;
}

const PING = "--";

export default function ConnectionNode({ isActive, onToggle }: ConnectionNodeProps) {
    const accent = isActive ? "blue" : "red";
    const statusColor = isActive ? "text-blue-400" : "text-red-500";

    return (
        <section className="px-10 flex flex-col items-start w-full">
            <div className="max-w-4xl flex flex-row items-center justify-start gap-12 py-6 w-full">

                <div className="flex flex-col">
                    <div className="flex items-center gap-3 mb-2">
                        <div className={`w-1.5 h-1.5 rounded-full ${isActive ? "bg-blue-500 shadow-[0_0_8px_#3b82f6]" : "bg-red-500 shadow-[0_0_8px_#ef4444]"}`} />
                        <span className="text-[9px] uppercase tracking-[0.3em] font-bold text-zinc-500">
                            Network Shield
                        </span>
                    </div>

                    <h2 className="text-6xl font-light tracking-tighter text-white leading-tight">
                        {isActive ? "SECURED" : "EXPOSED"}
                    </h2>

                    <p className={`text-[10px] mt-2 tracking-[0.1em] font-medium transition-colors duration-700 ${statusColor}/80`}>
                        {isActive ? "TOR CIRCUIT ESTABLISHED" : "ENCRYPTION OFF"}
                    </p>

                    <div className={`mt-6 border rounded-xl px-4 py-2 w-fit transition-all duration-500 ${isActive
                        ? "border-blue-500/50 bg-blue-500/5 shadow-[0_0_15px_rgba(59,130,246,0.1)]"
                        : "border-white/10 bg-white/5"
                        }`}>
                        <div className="flex items-center gap-3">
                            <span className="text-[10px] font-bold text-zinc-500 uppercase tracking-widest">Address</span>
                            <p className={`text-xs font-mono font-medium tracking-wider ${isActive ? "text-blue-400" : "text-zinc-400"}`}>
                                12.100.111.00
                            </p>
                        </div>
                    </div>
                </div>

                <div className="h-20 w-[1px] bg-zinc-800/80 mx-2" />

                <div className="flex flex-col items-center gap-4">
                    <button
                        onClick={onToggle}
                        className="group relative w-20 h-20 flex items-center justify-center cursor-pointer outline-none transition-all active:scale-95"
                    >
                        <div className={`absolute inset-0 rounded-full transition-opacity duration-1000 bg-blue-600/10 ${isActive ? "opacity-100" : "opacity-0"}`} />
                        <div className={`absolute inset-0 border rounded-full transition-all duration-700 ${isActive ? "border-blue-500/40 scale-110" : "border-zinc-700 hover:border-zinc-500"}`} />
                        <div className={`relative z-10 w-14 h-14 rounded-full flex items-center justify-center transition-all duration-500 border ${isActive ? "bg-white text-black border-white shadow-xl" : "bg-[#0d0d12] border-white/5 text-zinc-600"
                            }`}>
                            {isActive ? <ShieldCheck size={20} /> : <Power size={20} />}
                        </div>
                    </button>

                    <div className="flex items-center gap-1.5">
                        <Wifi size={12} className={isActive ? "text-blue-500" : "text-zinc-400"} />
                        <span className="text-[11px] font-mono text-zinc-400">{isActive ? PING : "--"}ms</span>
                    </div>
                </div>
            </div>

            <a
                href="https://github.com/MatthieuClaessens/BlackMist"
                target="_blank"
                rel="noopener noreferrer"
                className="fixed right-8 bottom-8 opacity-30 hover:opacity-100 transition-opacity duration-300 z-50 [--wails-draggable:no-drag]"
            >
                <img
                    src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original.svg"
                    alt="GitHub"
                    className="w-5 h-5 invert"
                />
            </a>
        </section>
    );
}