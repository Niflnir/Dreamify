import { Link } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import { Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

export default function App() {
  return (
    <SafeAreaView className='bg-primary h-full'>
      <View className='flex-1 items-center justify-center'>
        <Text className='text-3xl text-primary font-pbold'>Dreamify</Text>
        <Link className='text-primary font-pregular' href='/dreams'>Go To Dreams</Link>
        <StatusBar backgroundColor='#0F0F0F' />
      </View>
    </SafeAreaView>
  );
}

