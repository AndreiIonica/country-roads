import { RText } from '@/components/ui/Text';
import colors from '@/shared/colors';
import Constants from 'expo-constants';
import I18n from 'i18n-js';
import { FC } from 'react';
import { Pressable, StyleSheet, View } from 'react-native';

interface IProfileHeaderProps {
  onX?(): void;
}

const headerStyles = StyleSheet.create({
  headerContainer: {
    backgroundColor: colors.dark.accent,
    width: '100%',
    height: '13%',
  },
  settings: {},
  titleContainer: {
    top: Constants.statusBarHeight,
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
  },
});
export const ProfileHeader: FC<IProfileHeaderProps> = ({ onX }) => {
  return (
    <View style={headerStyles.headerContainer}>
      <View style={headerStyles.titleContainer}>
        <Pressable onPress={onX} style={headerStyles.settings} android_disableSound={true}>
          <RText text={I18n.t('settings')} variant="medium" />
        </Pressable>
        <Pressable onPress={onX} style={headerStyles.settings} android_disableSound={true}>
          <RText text={I18n.t('profile')} variant="semiBold" size="large" />
        </Pressable>
        <Pressable onPress={onX} style={headerStyles.settings} android_disableSound={true}>
          <RText text={I18n.t('logOut')} variant="medium" />
        </Pressable>
      </View>
    </View>
  );
};
