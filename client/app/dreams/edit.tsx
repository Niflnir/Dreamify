import { router } from "expo-router";
import { useState } from "react";
import { TextInput, TouchableOpacity, View, Text } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import Toast from "react-native-toast-message";

const EditDream = () => {
  const [title, setTitle] = useState('New Dream');
  const [body, setBody] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const onSaveHandler = async () => {
    try {
      setIsLoading(true);
      const response = await fetch('http://192.168.50.214:8080/posts', {
        method: 'POST',
        body: JSON.stringify({
          title: title,
          body: body
        })
      });
      const json = await response.json();
      console.log(json.data);
    } catch (err) {
      console.error(err);
    } finally {
      setIsLoading(false);
      Toast.show({
        type: 'success',
        text1: 'Dream saved successfully',
      })
      router.back();
    }
  }

  return (
    <SafeAreaView className='h-full bg-primary'>
      <View className='flex h-full flex-col px-3'>
        <TextInput
          className='rounded-lg p-2 font-pregular text-xl text-primary h-12'
          cursorColor='#532B88'
          autoFocus={true}
          placeholder='Add a title'
          placeholderTextColor='gray'
          value={title}
          onChangeText={(text) => {
            setTitle(text);
          }}
        />
        <TextInput
          className='h-80 rounded-lg p-2 font-pregular text-primary'
          cursorColor='#532B88'
          placeholder='What was your dream about?'
          placeholderTextColor='gray'
          textAlignVertical='top'
          multiline
          value={body}
          onChangeText={(text) => {
            setBody(text);
          }}
        />

        {/* Save button */}
        <View className="absolute bottom-4 right-6">
          <TouchableOpacity disabled={!title || !body} className="rounded-md bg-btnPrimary px-4 py-2" onPress={onSaveHandler}>
            <Text className='font-pregular text-primary'>Save</Text>
          </TouchableOpacity>
        </View>
      </View>
    </SafeAreaView>
  );
}

export default EditDream
