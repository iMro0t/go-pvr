/*
 *  Copyright (C) 2005-2018 Team Kodi
 *  This file is part of Kodi - https://kodi.tv
 *
 *  SPDX-License-Identifier: GPL-2.0-or-later
 *  See LICENSES/README.md for more information.
 */

#pragma once

#include "AddonBase.h"

#ifdef __cplusplus
extern "C"
{
#endif

  // ADDON_STATUS __declspec(dllexport) ADDON_Create(void *callbacks, void *props);
  // ADDON_STATUS __declspec(dllexport) ADDON_CreateEx(void *callbacks, const char *globalApiVersion, void *props);
  // void __declspec(dllexport) ADDON_Destroy();
  // ADDON_STATUS __declspec(dllexport) ADDON_GetStatus();
  // ADDON_STATUS __declspec(dllexport) ADDON_SetSetting(const char *settingName, const void *settingValue);
  //   __declspec(dllexport) const char *ADDON_GetTypeVersion(int type)
  //   {
  // #ifdef __cplusplus
  //     return kodi::addon::GetTypeVersion(type);
  // #endif
  //   }
  //   __declspec(dllexport) const char *ADDON_GetTypeMinVersion(int type)
  //   {
  // #ifdef __cplusplus
  //     return kodi::addon::GetTypeMinVersion(type);
  // #endif
  //   }

#ifdef __cplusplus
};
#endif
