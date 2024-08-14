import { View, Text, TouchableOpacity } from "react-native";
import { FC } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import BackIcon from "../../assets/icons/back-icon.svg"
import { router } from "expo-router";

interface HeaderProps { }

const Header: FC<HeaderProps> = () => {
    return (
        <SafeAreaView>
            <View className="bg-primary px-4 pt-4">
                <TouchableOpacity className="flex flex-row items-center gap-1.5" onPress={() => router.back()}>
                    <BackIcon width={20} height={20} />
                    <Text className="pt-0.5 font-pregular text-lg text-white">Back</Text>
                </TouchableOpacity>
            </View>
        </SafeAreaView>
    )
}

export default Header;
