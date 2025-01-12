import { StatusBar } from "expo-status-bar";
import { View, Text, TouchableOpacity, ScrollView, ActivityIndicator } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import WriteIcon from '../../assets/icons/write-icon.svg';
import { useEffect, useState } from "react";
import { Post } from "../../common/types";
import Dream from "../../common/components/Dream";
import { router } from "expo-router";
import { Colours } from "../../common/colours";

const Dreams = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await fetch('http://192.168.50.214:8080/posts');
        if (!response.ok) {
          throw new Error('Failed to retrieve posts')
        }

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
      } catch (err: any) {
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    }

    fetchPosts();
  }, [])

  if (isLoading) {
    return (
      <SafeAreaView className="h-full items-center justify-center bg-primary">
        <ActivityIndicator size="large" color={Colours.primary} />
        <Text className="mt-2 text-white">Loading...</Text>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className='h-full bg-primary'>
      <View className='flex flex-col'>
        <Text className='px-6 pb-4 pt-8 font-pbold text-3xl text-primary'>Dreams</Text>
        <ScrollView contentContainerStyle={{ paddingTop: 10, paddingBottom: 120 }}>
          <View className='size-full flex flex-col gap-4 px-6'>
            {
              posts.map((post) => (
                <Dream key={post.id} post={post} />
              ))
            }
          </View>
        </ScrollView>
      </View>

      {/* New dream button */}
      <View className='absolute bottom-6 right-6'>
        <TouchableOpacity className="rounded-full bg-btnPrimary p-4" onPress={() => router.push('/dreams/edit')}>
          <WriteIcon width={25} height={25} />
        </TouchableOpacity>
      </View>

      <StatusBar backgroundColor={Colours.background} />
    </SafeAreaView>
  );
}

export default Dreams
