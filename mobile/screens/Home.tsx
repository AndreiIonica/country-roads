import { RButton } from '@components/ui/Button';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import Colors from '@shared/colors';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ParamsList } from '../router/router';

type IHomeProps = NativeStackScreenProps<ParamsList, 'Home'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: Colors.dark.background,
  },
});

const HomeScreen: React.FC<IHomeProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <RButton
        text="Creare cont"
        style={{
          width: '91%',
        }}
      />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default HomeScreen;
