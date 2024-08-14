import { View, Text, Image } from "react-native";
import { Post } from "../types";
import { FC } from "react";

interface DreamProps {
    post: Post
}

const Dream: FC<DreamProps> = ({ post }) => {
    return (
        <View className="flex bg-gray-900 rounded-lg overflow-hidden my-2">
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
    )
}

export default Dream;
