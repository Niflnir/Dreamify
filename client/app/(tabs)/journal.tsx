import { StatusBar } from "expo-status-bar";
import { View, Text, TouchableOpacity, ScrollView, Image } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import EditIcon from '../../assets/icons/write-icon.svg';
import { useEffect, useState } from "react";

interface Post {
  id: number,
  title: string,
  body: string,
  dateCreated: string,
  imageUrl: string,
}

const Journal = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://192.168.50.214:8080/posts');
        const json = await response.json();
        const postsFromResponse: Post[] = [];
        json.data.forEach((data: any) => {
          const post: Post = {
            id: data.id,
            title: data.title,
            body: data.body,
            dateCreated: data.date_created,
            imageUrl: data.image_url
          }
          postsFromResponse.push(post);
        })
        setPosts(postsFromResponse);
      } catch (err) {
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    }

    fetchData();
  }, [])

  if (isLoading) {
    return <></>
  }

  return (
    <SafeAreaView className='h-full bg-primary'>
      <View className='flex flex-col'>
        <Text className='pt-8 pb-4 font-pbold text-3xl text-primary px-6'>Dreams</Text>
        <ScrollView contentContainerStyle={{ paddingBottom: 120 }}>
          <View className='flex flex-col size-full gap-4 px-6'>
            {
              posts.map((post) => (
                <View key={post.id} className="flex bg-gray-900 rounded-lg overflow-hidden">
                  {
                    post.imageUrl &&
                    <Image
                      className="absolute"
                      style={{ width: '100%', height: '100%', opacity: 0.3 }}
                      source={{ uri: post.imageUrl }}
                    />
                  }
                  <View className="flex gap-2 py-3 px-4">
                    <Text className='text-primary text-xl font-pbold'>{post.title}</Text>
                    <Text numberOfLines={5} className='text-gray-400 font-pregular text-justify'>{post.body}</Text>
                  </View>
                </View>
              ))
            }
          </View>
        </ScrollView>
      </View>

      {/* New dream button */}
      <View className='absolute bottom-4 right-6'>
        <TouchableOpacity className="rounded-full bg-btnPrimary p-4" onPress={() => console.log('pressed')}>
          <EditIcon width={25} height={25} />
        </TouchableOpacity>
      </View>

      <StatusBar backgroundColor='#0F0F0F' />
    </SafeAreaView>
  );
}

export default Journal
