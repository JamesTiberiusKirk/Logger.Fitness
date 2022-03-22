import { toastController } from '@ionic/vue';

export default async function notYetImplementedToast(){
  const toast = await toastController.create({
    message: "Not yet implemented",
    duration: 2000,
  });
  return toast.present();
}