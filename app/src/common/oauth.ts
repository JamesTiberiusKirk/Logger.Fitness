import { isPlatform, getPlatforms } from '@ionic/vue';

export function navigateToURL(url:string):void{
  console.log(getPlatforms());

  if (isPlatform("mobileweb") || isPlatform("desktop")){
    location.href = url;
  }
}
