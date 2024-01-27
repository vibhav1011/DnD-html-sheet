export default function Home(){
  return (
    <div>
      <h1>dnd-character-sheet</h1>
	  {Attribute("STR", "Strength")}
    </div>
  );
}

function Attribute(ATR, attribute) {
  return (
<div className="container max-w-20 relative">
  <ul>
    <li>

	<div class="group relative cursor-pointer py-2 container mx-auto bg-[#bbb] pb-[5px] rounded-[10px] flex flex-col justify-around items-center h-full z-1">
			<div class="absolute invisible top-1 group-hover:visible bg-[#bbb] text-black px-2 text-sm rounded-md">
			 <p class=" leading-2 text-gray-600 pt-1 ">{attribute}</p>
			 </div>
		   <span class="hover:cursor-pointer">{ATR}</span>
			<input name="Strengthscore" placeholder={10} className="container mx-auto rounded-[10px] flex flex-col justify-around items-center text-center py-1"/>
	 </div>

      <div className="container absolute transform translate-x-1/2 -translate-y-1/2 right-1/2 rounded-[20px] text-center border-[black] max-w-8 z-2">
        <input name="Strengthmod" placeholder={+0} className="container mx-auto rounded-[20px] flex flex-col justify-around items-center text-center border-[black] w-[35px] z-2"/>
      </div>

    </li>
  </ul>
</div>
  );
}

