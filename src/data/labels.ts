export const taskLabels: Record<string, string> = {
  chop_tree: '各プレイヤーが木を切ろう',
  curve: 'カーブ線路を25マス敷こう',
  dash: '50回ダッシュしよう',
  dont_catch_fire: '列車が燃えないようにしよう',
  hold_tools: '0:20以上続けて道具を持たないようにしよう',
  keep_crafter_busy: 'クラフトワゴンをxx秒間連続稼働させよう',
  kill_no_animals: '動物を殺さないようにしよう',
  let_train_burn: '列車を0:10秒間燃やそう',
  mine_rock: '各プレイヤーが岩を採掘しよう',
  one_player_axe: '一人だけが斧を持とう',
  one_player_track: '一人だけで線路を運ぼう',
  stack_items: '一箇所に物資を20個積み上げよう',
  straight: '直進線路をxxマス連続して敷こう',
  tracks_ahead: '線路を8マス以上先に伸ばそう',
};

export const extensionLabels: Record<string, string> = {
  charge: 'スーパーチャージ',
  drop: '自動ドロップ',
  front: 'フロントプル',
  ghost: 'ゴースト',
  hydro: 'ハイドロチャージ',
  magnet: 'マグネット吸引',
  pull: '貨物引き寄せ',
  push: '貨物押し出し',
  speed: 'スピードブースト',
  stack: 'スタックブースト',
};

export function translateTask(key: string): string {
  return taskLabels[key] ?? key;
}

export function translateExtension(key: string): string {
  return extensionLabels[key] ?? key;
}
