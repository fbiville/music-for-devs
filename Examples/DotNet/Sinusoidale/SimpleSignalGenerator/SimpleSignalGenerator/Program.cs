using NAudio.Wave;
using NAudio.Wave.SampleProviders;
using System;
using System.Threading;

namespace SimpleSignalGenerator
{
    internal class Program
    {
        static void Main(string[] args)
        {
            TestSignalGen();
            Console.ReadLine();
        }

        private static void TestSignalGen()
        {
            var A440Sin = new SignalGenerator()
            {
                Gain = 0.2,
                Frequency = 440,
                Type = SignalGeneratorType.Sin
            }.Take(TimeSpan.FromSeconds(5));


            using var wo1 = new WaveOutEvent();
            wo1.Init(A440Sin);
            wo1.Play();
            while (wo1.PlaybackState == PlaybackState.Playing)
            {
                Thread.Sleep(1000);
            }

        }
    }
}
