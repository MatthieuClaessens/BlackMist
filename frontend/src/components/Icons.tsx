const Logo = () => (
  <div className="relative w-8 h-8 flex items-center justify-center group cursor-pointer">
    <div className="absolute inset-0 border border-white rotate-45 transition-all duration-500 group-hover:bg-white group-hover:scale-110" />
    
    <div className="absolute w-full h-[2px] bg-black -rotate-45 z-10" />
    
    <div className="absolute inset-0 border border-white/20 rotate-45 scale-125 blur-sm" />
  </div>
);