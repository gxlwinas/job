import os
import pandas as pd
from pydub import AudioSegment
from datetime import datetime
import subprocess
import random

def convert_music_format(input_folder, output_folder):
	# 转换音乐文件格式为MP4
	for filename in os.listdir(input_folder):
		if filename.endswith('.m4a') or filename.endswith('.mp4'):
			input_file = os.path.join(input_folder, filename)
			output_file = os.path.join(output_folder, f'{os.path.splitext(filename)[0]}.mp4')

			# 使用FFmpeg进行格式转换
			cmd = ['ffmpeg', '-i', input_file, '-c:v', 'copy', '-c:a', 'aac', output_file]
			subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)

			print(f'转换完成: {input_file} -> {output_file}')

	print("音乐文件格式转换完成，已保存到指定的输出文件夹中。")


def convert_time_format(time_str):
	# 将时间字符串（格式为0:33:00）转换为毫秒
	time_parts = time_str.split(':')
	minutes, seconds, milliseconds = map(int, time_parts)
	total_milliseconds = (minutes * 60 + seconds) * 1000 + milliseconds
	return total_milliseconds


def load_and_edit_audio(df, music_folder, output_folder):
	output_audio_segments = []
	missing_audio = []

	for index, row in df.iterrows():
		song_name = row['歌名']
		start_time_str = str(row['开始'])  # 强制将时间列转换为字符串
		end_time_str = str(row['结束'])    # 强制将时间列转换为字符串

		# 解析开始时间和结束时间
		start_seconds = convert_time_format(start_time_str)
		end_seconds = convert_time_format(end_time_str)

		audio_file_path = os.path.join(music_folder, f'{song_name}.mp4')

		if not os.path.exists(audio_file_path):
			missing_audio.append(song_name)
			continue

		try:
			audio = AudioSegment.from_file(audio_file_path)
		except Exception as e:
			print(f"加载音频文件时出错：{str(e)}")
			continue

		# 计算开始和结束时间的毫秒
		start_time_ms = start_seconds
		end_time_ms = end_seconds
		print(start_seconds)
		print(end_seconds)
		trimmed_audio = audio[start_time_ms:end_time_ms]
		trimmed_audio = trimmed_audio.fade_in(2000).fade_out(2000)  # 添加渐进渐出效果

		output_path = os.path.join(output_folder, f'{song_name}.mp4')
		trimmed_audio.export(output_path, format='mp4')

		output_audio_segments.append(trimmed_audio)

		print(f"剪辑完成: {song_name}, 开始时间: {start_time_str}, 结束时间: {end_time_str}")

	print("音频剪辑完成。")

	if missing_audio:
		print("以下音乐文件缺失:")
		for missing_song in missing_audio:
			print(missing_song)

	return output_audio_segments


def concatenate_audio_segments(audio_segments, output_folder):
	combined_audio = AudioSegment.empty()
	for segment in audio_segments:
		combined_audio = combined_audio + segment

	# 输出拼接后的音频
	output_combined_path = os.path.join(output_folder, 'combined_output.mp4')
	combined_audio.export(output_combined_path, format='mp4')
	print(f"拼接后的音频保存在：{output_combined_path}")

def main():
	# 输入音乐文件夹路径
	music_folder = input("请输入音乐文件夹路径：")
	# music_folder = "./HYBE/rawMusicFile"
	# 输入Excel文件路径
	excel_file_path = input("请输入Excel文件路径：")
	# excel_file_path = "./HYBE/music.xlsx"
	# 输入输出根文件夹路径
	root_output_folder = input("请输入输出根文件夹路径：")
	# root_output_folder = "./"
	os.makedirs(root_output_folder, exist_ok=True)

	# 步骤1: 转换音乐文件格式
	converted_music_folder = os.path.join(root_output_folder, 'convertmusicfile')
	os.makedirs(converted_music_folder, exist_ok=True)
	convert_music_format(music_folder, converted_music_folder)

	# 读取 Excel 文件
	try:
		df = pd.read_excel(excel_file_path)
	except Exception as e:
		print(f"读取Excel文件时出错：{str(e)}")
		exit(1)

	# 步骤2: 加载音频文件并剪辑
	edited_music_folder = os.path.join(root_output_folder, 'editedmusicfile')
	os.makedirs(edited_music_folder, exist_ok=True)
	audio_segments = load_and_edit_audio(df, converted_music_folder, edited_music_folder)

	# 步骤3: 拼接音频
	output_music_folder = os.path.join(root_output_folder, 'outputmusicfile')
	os.makedirs(output_music_folder, exist_ok=True)
	concatenate_audio_segments(audio_segments, output_music_folder)

if __name__ == "__main__":
	main()