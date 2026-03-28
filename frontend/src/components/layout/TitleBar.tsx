import { CSSProperties } from "react";
import { X, Minus } from "lucide-react";
// @ts-ignore
import { Quit, WindowMinimise } from "../../../wailsjs/runtime/runtime";


interface WailsCSSProperties extends CSSProperties {
    "--wails-draggable"?: "drag" | "no-drag";
}

const draggableStyle: WailsCSSProperties = {
    "--wails-draggable": "drag"
};

export default function TitleBar() {
    return (
        <nav 
            style={draggableStyle} 
            className="flex-none h-14 flex items-center justify-between px-8 z-50 border-b border-white/2"
        >
            <div className="flex items-center gap-3 pointer-events-none">
                <h1 className="text-[18px] tracking-[0.4em] uppercase text-white/90">
                    BlackMist
                </h1>
            </div>
            
            <div className="flex items-center gap-2 [--wails-draggable:no-drag]">
                <button 
                    onClick={WindowMinimise} 
                    className="p-2 opacity-70 hover:opacity-100 transition-opacity cursor-pointer"
                >
                    <Minus size={14} />
                </button>
                <button 
                    onClick={Quit} 
                    className="p-2 opacity-70 hover:text-red-500 hover:opacity-100 transition-all cursor-pointer"
                >
                    <X size={14} />
                </button>
            </div>
        </nav>
    );
}