import { isPlatform, getPlatforms } from '@ionic/vue';

export function navigateToURL(url:string):void{
    console.log(getPlatforms());
    
    if (isPlatform("android")){
        console.log("web");
        console.log(url);
    }
}