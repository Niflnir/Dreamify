import { Tabs } from 'expo-router';
import BookIcon from '../../assets/icons/book-icon.svg';

const TabsLayout = () => {
  return (
    <>
      <Tabs screenOptions={{
        tabBarShowLabel: false,
        tabBarStyle: {
          backgroundColor: '#0F0F0F',
          borderTopColor: '#0F0F0F'
        }
      }}>
        <Tabs.Screen
          name='dreams'
          options={{
            title: 'Dreams',
            headerShown: false,
            tabBarIcon: () => {
              return (
                <BookIcon width={25} height={25} />
              )
            }
          }}
        />
      </Tabs>
    </>
  )
}

export default TabsLayout

